// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;



contract pancake {

    modifier checkDeadline(uint256 deadline) {
        require(block.timestamp<= deadline, 'Transaction too old');
        _;
    }
    
    function multicall(uint256 deadline, bytes[] calldata data)
        external
        payable
        checkDeadline(deadline)
        returns (bytes[] memory a)
    {
        return a;
         
    }

    
}


