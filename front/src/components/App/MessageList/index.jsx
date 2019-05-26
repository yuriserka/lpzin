import React from 'react';
import Mensagem from '../Message'

class MessageList extends React.Component {
    render() {
        return (
            <ul className= "message-list">
                {this.props.Mensagem.map(message => {
                    return (
                        <li key={message.id}>
                            <div>
                                message.senderId}
                            </div>
                            <div>
                                {message.text}
                            </div>
                        </li>
                    )
                })}
            </ul>
        )
    }
}