pragma solidity ^0.4.11;

contract test{
    function x() constant returns (uint) {
        uint a = 999560084690683640515566850045758101858937890625;
        uint b = 999780018149334509150625;
        return a / b;
    }
    function y() constant returns (uint) {
        uint a = 999560084690683640515566850045758101858937890625;
        uint b = 999780018149334509150625;
        uint c;
        assembly {
            c := div(a,b)
        }
        return c;
    }
}