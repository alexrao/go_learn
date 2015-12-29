#ifndef __GG_TCP_CLIENT_H__

#define GG_TCP_SERVER_IP    "121.46.0.200"
#define GG_TCP_SERVER_PORT   2001

enum
{
    GG_TCP_SEND_MODE_PLAY,
    GG_TCP_SEND_MODE_PAUSE,
};

int gg_tcp_data_send(unsigned char *data, int len);

int gg_tcp_send_funciton(int Key_mode);

#endif

