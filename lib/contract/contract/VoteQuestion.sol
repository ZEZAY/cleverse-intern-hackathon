/**
 *Submitted for verification at BscScan.com on 2022-07-25
 */

pragma solidity >=0.7.0 <0.9.0;

// SPDX-License-Identifier: MIT

contract VoteQuestion {
    string public questionName;
    string public questionDescription;
    address public contractAddress;
    address public contractCreator;
    string[] private choices;

    mapping(address => string) private mappingUserVote;
    mapping(string => uint256) private mappingChoiceData;

    struct TimestampData {
        uint256 startVoteAt;
        uint256 endVoteAt;
    }

    TimestampData timestampData;

    constructor(
        string memory _questionName,
        string memory _questionDescription,
        string[] memory _choices,
        uint256 _startQuestionAt,
        uint256 _duration
    ) {
        contractAddress = address(this);
        questionName = _questionName;
        questionDescription = _questionDescription;
        choices = _choices;
        timestampData = TimestampData(
            block.timestamp + _startQuestionAt,
            block.timestamp + _startQuestionAt + _duration
        );
    }

    function getAllChoices() public view returns (string[] memory) {
        return choices;
    }

    function getUserVoteData(address userAddress)
        public
        view
        returns (string memory)
    {
        return mappingUserVote[userAddress];
    }

    function _compareString(string memory a, string memory b)
        private
        pure
        returns (bool)
    {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
    }

    function vote(uint256 choiceId) public {
        require(
            _compareString(mappingUserVote[msg.sender], ""),
            "You already vote this question."
        );
        require(choiceId < choices.length, "Index out of range.");
        require(
            block.timestamp >= timestampData.startVoteAt &&
                block.timestamp <= timestampData.endVoteAt,
            "Voing period is end."
        );
        mappingUserVote[msg.sender] = choices[choiceId];
        mappingChoiceData[choices[choiceId]]++;
    }

    function getTimeStampData() public view returns (uint256, uint256) {
        return (timestampData.startVoteAt, timestampData.endVoteAt);
    }

    function getChoiceData(uint256 choiceId)
        public
        view
        returns (uint256, string memory)
    {
        require(choiceId < choices.length, "Index out of range.");
        return (mappingChoiceData[choices[choiceId]], choices[choiceId]);
    }
}
