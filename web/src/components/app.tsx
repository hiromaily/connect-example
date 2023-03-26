import { useState } from "react";
import { GreetService } from "../gen/greet/v1/greet_connectweb";
import {useClient} from "../libs/client";

// https://github.com/depot/sdk-node/blob/7c1bc12a70b415111abd25fd7695ab37f03685b8/src/index.ts

const App = () => {
    const [inputValue, setInputValue] = useState("");
    // @ts-ignore
    const client = useClient(GreetService);

    return <>
        <form onSubmit={async (e) => {
            e.preventDefault();
            // @ts-ignore
            const res = await client.greet({
                name: inputValue,
            });
            console.log(res);
        }}>
            <input value={inputValue} onChange={e => setInputValue(e.target.value)} />
            <button type="submit">Send</button>
        </form>
    </>;
}

export default App