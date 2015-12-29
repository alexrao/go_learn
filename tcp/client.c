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

#define    MAXLINE        1024
#define GG_TCP_SERVER_IP    "121.46.0.200"
//#define GG_TCP_SERVER_IP    "192.168.1.105"
#define GG_TCP_SERVER_PORT   2001

int send_test(void)
{
    struct sockaddr_in    serv_addr;
    int                    sock_id;
    int                    send_len;
    int                    i_ret;
    char string[4];
    int len = 3;

    string[0] = 0xFF;
    string[1] = 0xFF;
    string[2] = 0xFF;
    string[3] = 0;

    /* create the socket commented by guoqingbo*/
    if ((sock_id = socket(AF_INET, SOCK_STREAM, 0)) < 0)
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

    //while(1)
    {
        /* transported the file commented by guoqingbo*/
        send_len = send(sock_id, string, len, MSG_WAITALL);
        printf("Send Data len :%d\n", send_len);
        if (send_len < 0 )
        {
            perror("Send file failed\n");
            exit(0);
        }
        sleep(3);
    }

    close(sock_id);
    printf("Send Success\n");
    return 0;
}

void main(void)
{
    while(1)
    {
        send_test();
        sleep(3);
    }
}
