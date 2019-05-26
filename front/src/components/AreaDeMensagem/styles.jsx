import styled from 'styled-components';

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

export {PainelEnviarMensagem, InputMensagem};
