export interface Todo {
  id:  string;
  description:  string;
  isDone: boolean;
}

export function createTodo(params: Partial<Todo>) {
  return {

  } as Todo;
}
