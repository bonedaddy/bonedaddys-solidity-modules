pragma solidity >=0.4.24 <=0.8.0;

// used to measure the gas costs of initializing a struct in two forms
// repeated assignments to struct vs singular assignment

contract StructAssignTest {
    struct numbers {
        uint256 a;
        uint256 b;
        uint256 c;
    }

    mapping (uint256 => numbers) public structs;

    function setMany() public returns (bool) {
        structs[block.number].a = 1;
        structs[block.number].b = 2;
        structs[block.number].c = 3;
        return true;
    }

    // cheaper than setMany by about 150 gas
    function setSingle() public returns (bool) {
        structs[block.number] = numbers({
            a: 1,
            b: 2,
            c: 3
        });
    }
}