import styled from 'styled-components';

const PainelEnviarMensagem = styled.div`
    display: flex;
    flex-direction: row;
    width: 51%;
    height: 60px;
    position: relative;
    z-index: 1;
    position: fixed;
    bottom: 9px;
`;

const PainelDigitarMensagem = styled.div`
    background: whitesmoke;
    width: 100%;
    padding: 10px 20px;
`;

const InputMensagem = styled.input`
    outline: none;
    width: 100%;
    border-radius: 10px;
    height: 36px;
`;

export {PainelEnviarMensagem, InputMensagem, PainelDigitarMensagem};
