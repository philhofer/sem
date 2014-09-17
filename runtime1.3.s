// +build go1.3

#define NOSPLIT 4

#ifdef GOARCH_arm
#define JMP B
#endif

// semacquire(sem *uint32)
TEXT ·semacquire(SB),NOSPLIT,$0-0
	JMP runtime·semacquire(SB)



// Semrelease(sem *uint32)
TEXT ·semrelease(SB),NOSPLIT,$0-0
	JMP runtime·semrelease(SB)


