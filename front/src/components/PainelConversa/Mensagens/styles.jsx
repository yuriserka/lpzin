import styled from 'styled-components';

const MensagemDiv = styled.div`
    background: white;
    width: 50%;
    float: left;
    border-radius: 10px;
    margin: auto;
    border: 1px solid rgba(0, 0, 0, 0);
    -moz-background-clip: border;
    -webkit-background-clip: border;
    background-clip: border-box;
                    
    -moz-background-clip: padding;
    -webkit-background-clip: padding;
    background-clip: padding-box;
                    
    -moz-background-clip: content;
    -webkit-background-clip: content;
    background-clip: content-box;
`;

const Sender = styled.div`
    padding-left: 5px;
`;

const MessageContent = styled.div`
    padding-left: 5px;
    margin-bottom: 3px;
    overflow-wrap: break-word;
    word-wrap: break-word;
    hyphens: auto;
    @import url('https://fonts.googleapis.com/css?family=Quicksand');
    font-family: 'Quicksand';
    font-size: 14px;
`;

export {MensagemDiv, Sender, MessageContent};
