.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $0, %eax
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

