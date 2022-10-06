import { FormEvent, useState } from "react";
import "./App.css";
import { useGetTodosQuery } from "./graphql/generated";

import { useQuery, useMutation, useSubscription, gql } from "urql";

function App() {
  const [text, setText] = useState("");
  const [done, setDone] = useState(false);

  const [getTodosResult, rexececuteQuery] = useGetTodosQuery();

  const [createTodoResult, createTodo] = useMutation(gql`
    mutation ($input: CreateTodoInput!) {
      createTodo(input: $input) {
        id
        text
        done
      }
    }
  `);

  const [subscriptionResult] = useSubscription({
    query: gql`
      subscription {
        todoCreated {
          id
          text
          done
        }
      }
    `,
  });

  const handleSubmit = (e: FormEvent) => {
    //prevent the page from refreshing
    e.preventDefault();

    createTodo({
      input: {
        text,
        done,
      },
    });

    // reset form values
    setText("");
    setDone(false);
  };

  return (
    <div className="App">
      <form
        onSubmit={handleSubmit}
        style={{ display: "flex", flexDirection: "column" }}
      >
        <input
          type="text"
          placeholder="todo text"
          value={text}
          onChange={(e) => setText(e.currentTarget.value)}
        />
        <div>
          <label>done</label>
          <input
            type="checkbox"
            checked={done}
            onChange={(e) => setDone(e.currentTarget.checked)}
          />
        </div>
        <input type="submit" value="add" />
      </form>
      <table className="styled-table">
        <thead>
          <tr>
            <th>id</th>
            <th>text</th>
            <th>done</th>
          </tr>
        </thead>
        <tbody>
          {getTodosResult.data?.getTodos.map((todo) => (
            <tr key={todo?.id}>
              <td>{todo?.id}</td>
              <td>{todo?.text}</td>
              <td>{todo?.done ? "Yes" : "No"}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <div>
        <h3>Latest created Todo</h3>
        {JSON.stringify(subscriptionResult.data?.todoCreated)}
      </div>
    </div>
  );
}

export default App;
