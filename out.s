.data
.balign 8
lol.0:
	.ascii "\"Hello, World!\"\n"
	.byte 0
/* end data */

.text
lol:
	pushq %rbp
	movq %rsp, %rbp
	leaq lol.0(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $0, %eax
	leave
	ret
.type lol, @function
.size lol, .-lol
/* end function lol */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	callq lol
	movl $0, %eax
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
