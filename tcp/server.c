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
#include <sys/errno.h>
#include <signal.h>

#define    MAXLINE        1024
#define GG_TCP_SERVER_PORT   2001

int main(void)
{
    int listenfd, connectfd;
    struct sockaddr_in    server;
    struct sockaddr_in    client;
    socklen_t addrlen;

        if((listenfd = socket(AF_INET, SOCK_STREAM, 0)) == -1)
        {
            perror("Creating  socket failed.");
            exit(1);
        }
    int opt =SO_REUSEADDR;
    setsockopt(listenfd,SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));
    bzero(&server,sizeof(server));
    server.sin_family=AF_INET;
    server.sin_port=htons(GG_TCP_SERVER_PORT);
    server.sin_addr.s_addr= htonl (INADDR_ANY);
    if(bind(listenfd, (struct sockaddr *)&server, sizeof(server)) == -1)
    {
        perror("Binderror.");
        exit(1);
    }

    if(listen(listenfd, 1)== -1)
    {  /* calls listen() */
        perror("listen()error\n");
        exit(1);
    }
    int recv_len, i;
    unsigned char recv_buf[1024];
    unsigned char send_buf[1024];

    while(1)
    {
        addrlen =sizeof(client);
        if((connectfd = accept(listenfd,(struct sockaddr*)&client,&addrlen)) < 0)
        {
            perror("accept()error\n");
           if (errno == EINTR)
           {
               perror("Cause it\n");
               continue;
           }
           else
           {
               break;
           }
        }

        recv_len = recv(connectfd, recv_buf, sizeof(recv_buf), 0);

        if(recv_len > 0)
        {
            for(i=0; i<recv_len; i++)
            {
                printf("0x%02x, ", recv_buf[i]);
            }
            printf("\n");
        }

        snprintf(send_buf, sizeof(send_buf), "Get Data Ok\n");
        send(connectfd, send_buf, strlen(send_buf), 0);
        close(connectfd);
    }

    close(listenfd);
    return 0;
}
