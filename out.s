.text
abc:
	pushq %rbp
	movq %rsp, %rbp
	movl $12, %eax
	leave
	ret
.type abc, @function
.size abc, .-abc
/* end function abc */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	callq abc
	movl %eax, %esi
	leaq fmt(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $0, %eax
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.data
.balign 8
fmt:
	.ascii "abc: %d\n"
	.byte 0
/* end data */

.section .note.GNU-stack,"",@progbits
