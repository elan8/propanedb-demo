import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ID } from '@datorama/akita';
import { tap } from 'rxjs/operators';
import { createTodo, TodoItem, TodoItems } from './todo.model';
import { TodoStore } from './todo.store';

@Injectable({ providedIn: 'root' })
export class TodoService {

  constructor(private todoStore: TodoStore, private http: HttpClient) {
  }


  getAll() {
    // return this.http.get<TodoItem[]>('http://localhost:8080/v1/todo').pipe(tap(response => {
    //   let todos = response as TodoItem[];
    //   this.todoStore.update({
    //     todos: todos
    //   });
    //   //this.todoStore.set(entities);
    // }));

    return this.http.get('http://localhost:8080/v1/todo', { observe: 'response' }).pipe(tap(response => {
      //console.log( JSON.stringify(response) )

     let todoItems = response?.body as TodoItems;
     let todos = todoItems.items;
      this.todoStore.update({
        todos: todos
      });

      console.log( JSON.stringify(todos) )


      //  //this.todos= data
      // let todos = response as TodoItem[];
      // this.todoStore.update({
      //   todos: todos
      // });
      //this.todoStore.set(entities);
    }));

  }

  add(todo: TodoItem) {

    return this.http.post<TodoItem[]>('http://localhost:8080/v1/todo/create', todo).pipe(tap(entities => {
      //this.todoStore.set(entities);
    }));

  }

  update(todo: Partial<TodoItem>) {
    return this.http.post<TodoItem[]>('http://localhost:8080/v1/todo/update', todo).pipe(tap(entities => {
      //this.todoStore.set(entities);
    }));
    //this.todoStore.update(id, todo);
  }

  remove(id: ID) {

    return this.http.delete<TodoItem>('http://localhost:8080/v1/todo/delete/'+id).pipe(tap(entities => {
      //this.todoStore.set(entities);
    }));

    //this.todoStore.remove(id);
  }

}
