


#include "wt_demo.h"




static const char *home;


int
main(int argc, char *argv[])
{
    WT_CONNECTION *conn;
    WT_CURSOR *cursor;
    WT_SESSION *session;

    home = "/tmp/wt";
    const char *key, *value;
    int ret;
    error_check(wiredtiger_open(home, NULL, "create", &conn));

    error_check(conn->open_session(conn, NULL, NULL, &session));

    error_check(session->create(session, "table:demo", "key_format=S,value_format=S"));
    error_check(session->open_cursor(session, "table:demo", NULL, NULL, &cursor));
    
    cursor->set_key(cursor, "foo");
    cursor->set_value(cursor, "bar");
    error_check(cursor->insert(cursor));

    cursor->set_key(cursor, "foo");
    cursor->set_value(cursor, "newbar");
    error_check(cursor->update(cursor));

    cursor->set_key(cursor, "go");
    cursor->set_value(cursor, "val");
    error_check(cursor->insert(cursor));

    error_check(cursor->reset(cursor));

    while ((ret = cursor->next(cursor)) == 0) {
        error_check(cursor->get_key(cursor, &key));
        error_check(cursor->get_value(cursor, &value));

        printf("Got record: %s, %s\n", key, value);
    }

    error_check(conn->close(conn, NULL));
    return 0;
}

/*

git clone git://github.com/wiredtiger/wiredtiger.git

apt install -y autoconf automake libtool
cd wiredtiger
sh autogen.sh
/configure && make
make install

gcc -g  wt_demo.c /usr/local/lib/libwiredtiger.a -lpthread -ldl -o wtdemo
*/
