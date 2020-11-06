pragma solidity >=0.4.24 <=0.8.0;


contract Administration {

    address public owner;
    address public admin;
    bool    public frozen;

    mapping (address => bool) public moderators;

    event AdminSet(address _admin);
    event ModeratorSet(address _moderator);
    event ModeratorUnset(address _moderator);
    event OwnershipTransferred(address _previousOwner, address _newOwner);

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == owner || msg.sender == admin);
        _;
    }

    modifier onlyPrivileged() {
        require(msg.sender == owner || msg.sender == admin || moderators[msg.sender] == true);
        _;
    }

    constructor() {
        owner = msg.sender;
        admin = msg.sender;
    }

    function setAdmin(
        address _newAdmin
    )
        public
        onlyOwner
        returns (bool)
    {
        require(_newAdmin != admin);
        admin = _newAdmin;
        emit AdminSet(_newAdmin);
        return true;
    }

    function setModerator(
        address _moderator
    )
        public
        onlyAdmin
        returns (bool)
    {
        moderators[_moderator] = true;
        emit ModeratorSet(_moderator);
        return true;
    }

    function unsetModerator(
        address _moderator
    )
        public
        onlyAdmin
        returns (bool)
    {
        delete moderators[_moderator];
        emit ModeratorUnset(_moderator);
        return true;
    }

    function transferOwnership(
        address _newOwner
    )
        public
        onlyOwner
        returns (bool)
    {
        require(_newOwner != owner && _newOwner != address(0));
        owner = _newOwner;
        emit OwnershipTransferred(msg.sender, _newOwner);
        return true;
    }
}
