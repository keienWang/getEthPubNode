// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.0;

contract Ma_p {
    uint256 public i = 1;
    mapping(uint256 => address) public contracts;
    function Set(address c)public{
        contracts[i] = c;
        i++;
    }
}