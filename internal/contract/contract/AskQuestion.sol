/**
 *Submitted for verification at BscScan.com on 2022-07-25
 */

pragma solidity >=0.7.0 <0.9.0;

// SPDX-License-Identifier: MIT

contract AskQuestion {
    string public questionName;
    string public questionDescription;
    address public contractAddress;
    address public contractCreator;
    string[] private questionAnswers;

    mapping(address => string) private mappingUserAnswer;

    struct TimestampData {
        uint256 startQuestionAt;
        uint256 endQuestionAt;
    }

    TimestampData timestampData;

    constructor(
        string memory _questionName,
        string memory _questionDescription,
        uint256 _startQuestionAt,
        uint256 _duration
    ) {
        contractAddress = address(this);
        questionName = _questionName;
        questionDescription = _questionDescription;
        timestampData = TimestampData(
            block.timestamp + _startQuestionAt,
            block.timestamp + _startQuestionAt + _duration
        );
    }

    function getAllAnswer() public view returns (string[] memory) {
        return questionAnswers;
    }

    function getUserAnswerData(address userAddress)
        public
        view
        returns (string memory)
    {
        return mappingUserAnswer[userAddress];
    }

    function _compareString(string memory a, string memory b)
        private
        pure
        returns (bool)
    {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
    }

    function answer(string memory userAnswer) public {
        require(
            _compareString(mappingUserAnswer[msg.sender], ""),
            "You already answer this question."
        );
        require(
            block.timestamp >= timestampData.startQuestionAt &&
                block.timestamp <= timestampData.endQuestionAt,
            "Question period is end."
        );
        mappingUserAnswer[msg.sender] = userAnswer;
        questionAnswers.push(userAnswer);
    }

    function getTimeStampData() public view returns (uint256, uint256) {
        return (timestampData.startQuestionAt, timestampData.endQuestionAt);
    }
}
