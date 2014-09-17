// +build go1.4

#define NOSPLIT 4

TEXT ·syncsemacquire(SB),NOSPLIT,$0-0
	JMP runtime·syncsemacquire(SB)



TEXT ·syncsemrelease(SB),NOSPLIT,$0-0
	JMP runtime·syncsemrelease(SB)


