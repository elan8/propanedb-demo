export interface TodoItem {
  id:  string;
  title:  string;
  completed: boolean;
}

export function createTodo(title: string) {

  return {
    id: "",
    title,
    completed: false
  } as TodoItem;
}