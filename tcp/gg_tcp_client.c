#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include "s100_global.h"
#include "dprintf.h"
#include "gg_tcp_client.h"

int gg_tcp_data_send(unsigned char *data, int len)
{
    struct sockaddr_in    serv_addr;
    int                    sock_id;
    int                    send_len;
    int                    i_ret;

    /* create the socket commented by guoqingbo*/
    if ((sock_id = socket(AF_INET,SOCK_STREAM,0)) < 0)
    {
        perror("Create socket failed\n");
        exit(0);
    }

    memset(&serv_addr, 0, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(GG_TCP_SERVER_PORT);
    inet_pton(AF_INET, GG_TCP_SERVER_IP, &serv_addr.sin_addr);

    /* connect the server commented by guoqingbo*/
    i_ret = connect(sock_id, (struct sockaddr *)&serv_addr, sizeof(struct sockaddr));
    if (-1 == i_ret)
    {
        perror("Connect socket failed\n");
        return -1;
    }

    /* transported the file commented by guoqingbo*/
    send_len = send(sock_id, data, len, MSG_WAITALL);
    if (send_len < 0 )
    {
        perror("Send file failed\n");
        exit(0);
    }
    else
    {
        sleep(1);
    }

    close(sock_id);
    printf("Send Finish\n");
    return 0;
}

int gg_tcp_send_funciton(int Key_mode)
{
    int len = 0;
    unsigned char data[1024];

    bzero(data, sizeof(data));
    if(Key_mode == GG_TCP_SEND_MODE_PLAY)
    {
        data[0] = 0xFF;
        data[1] = 0xFF;
        data[2] = 0xFF;
        len = 3;
        DPRINTF(1, "Play");
    }
    else if(Key_mode == GG_TCP_SEND_MODE_PAUSE)
    {
        data[0] = 0;
        data[1] = 0;
        data[2] = 0;
        len = 3;
        DPRINTF(1, "Pause");
    }
    else
    {
        DPRINTF(1, "Not Correct Mode");
    }

    if(len > 0)
    {
        if(gg_tcp_data_send(data, len) < 0)
        {
            DPRINTF(4, "===============================");
            DPRINTF(4, "===============================");
            DPRINTF(2, "GG TCP Data Send Failed");
            DPRINTF(4, "===============================");
            DPRINTF(4, "===============================");
            return -1;
        }
        else
        {
            DPRINTF(4, "===============================");
            DPRINTF(4, "===============================");
            DPRINTF(2, "GG TCP Data Send OK");
            DPRINTF(4, "===============================");
            DPRINTF(4, "===============================");
            return 1;
        }
    }

    return 0;
}

