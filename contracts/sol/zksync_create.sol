pragma solidity ^0.8.0;


contract ContractDeployer {
   

    function create(
        bytes32 _salt,
        bytes32 _bytecodeHash,
        bytes calldata _input
    ) external payable  returns (address) {
        return address(this);
        // return createAccount(_salt, _bytecodeHash, _input, AccountAbstractionVersion.None);
    }

      function withdraw(address _l1Receiver) external payable {
        uint256 amount = msg.value;

    }

}
