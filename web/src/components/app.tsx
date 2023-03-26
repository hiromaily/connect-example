import {FormEvent, useState} from "react";
import { GreetService } from "../gen/greet/v1/greet_connect";
import { useClient } from "../libs/client";

// https://github.com/depot/sdk-node/blob/7c1bc12a70b415111abd25fd7695ab37f03685b8/src/index.ts

const App = () => {
    const [inputValue, setInputValue] = useState("");
    const [response, setResponse] = useState("");
    // @ts-ignore
    const client = useClient(GreetService);

    const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        // @ts-ignore
        const res = await client.greet({
            name: inputValue,
        });
        setResponse(res.greeting)
        console.log(res);
    }

    return <>
        <p>Say, something</p>
        <form onSubmit={onSubmit}>
            <input value={inputValue} onChange={e => setInputValue(e.target.value)} />
            <button type="submit">Send</button>
        </form>
        <p>Response:<span>{response}</span></p>
    </>;
}

export default App