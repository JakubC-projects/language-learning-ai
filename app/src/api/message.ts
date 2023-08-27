import { Ref, ref } from "vue"
import { Message } from "../types/message"

const textDecoder = new TextDecoder()
export async function sendPrompt(messages: Message[]):Promise<Ref<Message>> {
    const res = await fetch("http://localhost:8081/message", {
        method: "POST",
        body: JSON.stringify({
            messages
        })
    })
    if(!res.body) throw Error("missing body")
    const result = ref<Message>({content: '', role: 'assistant'})

    const reader = res.body?.getReader()
    fillMessage(reader, result)

    return result
}

async function fillMessage(reader: ReadableStreamDefaultReader<Uint8Array>, message: Ref<Message>) {
    while(true) {
        const chunk =  await reader.read()
        if (chunk.value) {
            var string = textDecoder.decode(chunk.value);
            message.value.content += string;
            console.log(string)
        }
        if(chunk.done) {
            break;
        }
    }
}

