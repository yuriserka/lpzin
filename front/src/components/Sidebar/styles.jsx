import styled from 'styled-components';

const SidebarDiv = styled.div`
    width: 31%;
    height: 100%;
    background: white;
    z-index: 0;
    float: left;
    position: relative;
`;

const ConversasDivStyle = styled.div`
  margin-top: 5px;
`;

const ConversaDiv = styled.div`
  &:hover {
    background: lightgray;
  }
`;

const SeparadorConversa = styled.div`
  border: 0;
  height: 1px;
  background-image: linear-gradient(to right, rgba(0, 0, 0, 0),
    rgba(0, 0, 0, 0.75), rgba(0, 0, 0, 0));
`;

export {SidebarDiv, ConversasDivStyle, ConversaDiv, SeparadorConversa};
