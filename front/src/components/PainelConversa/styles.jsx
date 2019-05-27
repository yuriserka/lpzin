import styled from 'styled-components';

const PainelConversaDiv = styled.div`
    background: violet;
    height: 100%;
    width: 69%;
    float: left;
    position: relative;
`;

const PainelEnviarMensagem = styled.form`
    display: flex;
    flex-direction: row;
    max-width: 100%;
    min-height: 62px;
    position: relative;
    z-index: 2;
`;

const InputMensagem = styled.input`
    line-height: 20px;
    width: 100%;
`;

export {PainelConversaDiv, PainelEnviarMensagem, InputMensagem};
