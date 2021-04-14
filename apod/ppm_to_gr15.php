#!/usr/bin/php
<?php
/* Skip header; assuming:
P6
160 192
255
*/

for ($i = 0; $i < 3; $i++) {
  fgets(STDIN);
}

/* Generate array to store image, so we can go over it a few times */
$px = array();
for ($y = 0; $y < 192; $y++) {
  $px[$y] = array();
  for ($x = 0; $x < 160; $x++) {
    $px[$y][$x] = array(0, 0, 0);
  }
}

/* Space to store colors: */

$colors = array();

/* Load the image */
for ($y = 0; $y < 192; $y++) {
  for ($x = 0; $x < 160; $x++) {
    $r = ord(fgetc(STDIN));
    $g = ord(fgetc(STDIN));
    $b = ord(fgetc(STDIN));
    $px[$y][$x][0] = $r;
    $px[$y][$x][1] = $g;
    $px[$y][$x][2] = $b;

    $c = sprintf("%02x%02x%02x", clamp($r), clamp($g), clamp($b));

    if (!in_array($c, $colors)) {
      $colors[] = $c;
    }
  }
}

$palette = array();
$idx = 0;
foreach ($colors as $c) {
  $palette[$c] = $idx++;
}

$b = array(0, 0, 0, 0);

for ($y = 0; $y < 192; $y++) {
  for ($x = 0; $x < 160; $x += 4) {
    for ($i = 0; $i < 4; $i++) {
      $color = sprintf("%02x%02x%02x",
        clamp($px[$y][$x + $i][0]),
        clamp($px[$y][$x + $i][1]),
        clamp($px[$y][$x + $i][2])
      );

      $b[$i] = $palette[$color];
    }

    $c = chr(($b[0] * 64) + ($b[1] * 16) + ($b[2] * 4) + $b[3]);
    fwrite(STDOUT, $c, 1);
  }
}

fprintf(STDERR, "palette = \n%s\n\n", print_r($palette, true));

/* Determine the Atari colors utilize, so we can send bytes for the four
   color palette entries */

$fi = fopen("atari256.ppm", "r");
/* Skip header; assuming:
P6
160 192
255
*/

for ($i = 0; $i < 3; $i++) {
  fgets($fi);
}

$atari_colors = array();
for ($i = 0; $i < 256; $i++) {
  $r = ord(fgetc($fi));
  $g = ord(fgetc($fi));
  $b = ord(fgetc($fi));

  $c = sprintf("%02x%02x%02x", clamp($r), clamp($g), clamp($b));
  $atari_colors[$c] = $i;
}
fclose($fi);

fprintf(STDERR, "atari_colors = \n%s\n\n", print_r($atari_colors, true));

foreach ($palette as $rgb => $_) {
  if (!array_key_exists($rgb, $atari_colors)) {
    echo "color $rgb doesn't exist\n";
    fwrite(STDOUT, chr(0), 1);
  } else {
    $c = chr($atari_colors[$rgb]);
    fwrite(STDOUT, $c, 1);
  }
}

function clamp($x) {
  return (floor($x / 16) * 16);
}

