import { Chat } from "../types/chat";

export async function getChats(): Promise<Chat[]> {
    return []
}

export async function createChat(name: string): Promise<Chat> {
    return {
        id: "",
        name,
        userId: "",
    }
}