import styled from 'styled-components';

const ConversaHeaderDiv = styled.div`
    color: white;
    background: #921840;
    height: 48px;
    margin: 0;
    margin-top: -48px;
    float:left;
    width: 69%;
    &:hover {
        background: #5a0e27;
    }
`;

const ConversaInfoDiv = styled.div`
    padding: 16px 20px;
`;

export {ConversaHeaderDiv, ConversaInfoDiv};
