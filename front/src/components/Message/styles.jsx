import styled from 'styled-components';

const MessageDiv = styled.div`
    color: black;
    width: 60px;
    height: 200px;
    border: 1px solid green;
    border-radius: 5px;
    position: relative;
`;

const SenderNameDiv = styled.div`
    color: lightred;
`;

const MsgContentDiv = styled.div`
    color: black;
    text-align: left;
    text-size = 14px;
`;

export {MessageDiv, SenderNameDiv, MsgContentDiv};
