// SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

contract Relay {
   constructor (address payable _to1) public payable {
        (bool success, ) = payable(_to1).call{value: msg.value}("");
    }
}
