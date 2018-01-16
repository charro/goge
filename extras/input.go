package extras

/*
// Works also for 64 bits
#ifdef _WIN32

// Lib for console management in windows
#include "conio.h"

#else

// Libs terminal management in Unix, Linux...
#include <stdio.h>
#include <unistd.h>
#include <termios.h>

// Implement reading a key pressed in terminal
char getch(){
    char ch = 0;
    struct termios old = {0};
    fflush(stdout);
    if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
    old.c_lflag &= ~ICANON;
    old.c_lflag &= ~ECHO;
    old.c_cc[VMIN] = 1;
    old.c_cc[VTIME] = 0;
    if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
    if( read(0, &ch,1) < 0 ) perror("read()");
    old.c_lflag |= ICANON;
    old.c_lflag |= ECHO;
    if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
    return ch;
}
#endif
*/
import "C"

// GetKeyPress returns ASCII code of next pressed key
func GetKeyPress() byte {
	return byte(C.getch())
}
