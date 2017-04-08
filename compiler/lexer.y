/* Declarations */

%{
package compiler

%}

%union {
    b int  // byte
}

// %tokens without a type (keyword) don't need a type
%token INC_PTR, DEC_PTR, OTHER

// only tokens with a value need a type
// %token <type> <TOKEN>
%token <b> INC_BYTE, DEC_BYTE, READ_BYTE, WRITE_BYTE


%start expr

%%  /* Grammar rules below */

expr        : loop
            | expr loop
            | command
            | expr command
            ;

loop        : '[' expr ']'
            | '[' ']' // empty loop, pointless but valid
            ;

command     : INC_BYTE
            | DEC_BYTE
            | READ_BYTE
            | WRITE_BYTE
            | INC_PTR
            | DEC_PTR
            | OTHER
            ;

%%  /* Program */
