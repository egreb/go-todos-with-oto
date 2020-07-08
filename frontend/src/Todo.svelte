<script>
  import { fade } from "svelte/transition";
  import { TodosService } from "./oto.gen.js";
  const todoService = new TodosService();
  export let id;
  export let task;
  export let completed;
  export let mark;
  export let remove;

  const handleUpdate = () => {
    todoService
      .update({
        todo: {
          id,
          task,
          completed: !completed
        }
      })
      .then(result => {
        const { todo } = result;
        task = todo.task;
        completed = todo.completed;

        mark(todo);
      })
      .catch(err => {
        console.log("todo err", err);
      });
  };

  const handleDelete = () => {
    todoService
      .delete({
        todo: {
          id,
          task,
          completed
        }
      })
      .then(res => remove({ id }))
      .catch(err => console.error);
  };
</script>

<style>
  li {
    border: 1px solid #999;
    box-shadow: 1 1 1 1 #333;
    width: 100%;
    border-radius: 5px;
    margin-top: 5px;
    list-style: none;
    padding: 0.5rem;
    font-size: 1.5rem;
    display: flex;
    align-items: center;
  }
  li:first-child {
    margin-top: 0;
  }

  span.done {
    text-decoration: line-through;
  }

  input[type="checkbox"] {
    margin: 0;
    padding: 0.5rem;
  }

  span {
    padding: 0 0.5rem;
  }
  div {
    margin-left: auto;
  }
  button {
    margin-left: auto;
    border: none;
    background: none;
    padding: 0;
    margin: 0;
    color: red;
    align-self: flex-end;
    cursor: pointer;
    font-size: 0.95rem;
    display: flex;
    align-items: center;
    padding: 0.5rem;
  }

  button:hover {
    background-color: #f7f7f7;
    border: 1px solid #999;
    border-radius: 5px;
  }
</style>

<li transition:fade>
  <input type="checkbox" bind:checked={completed} on:click={handleUpdate} />
  <span class:done={completed}>{task}</span>
  <div>
    <button on:click|preventDefault={handleDelete}>trash</button>
  </div>
</li>
