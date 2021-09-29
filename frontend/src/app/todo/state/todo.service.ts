import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ID } from '@datorama/akita';
import { tap } from 'rxjs/operators';
import { createTodo, TodoItem } from './todo.model';
import { TodoStore } from './todo.store';

@Injectable({ providedIn: 'root' })
export class TodoService {

  constructor(private todoStore: TodoStore, private http: HttpClient) {
  }


  get() {
    return this.http.get<TodoItem[]>('http://localhost:8080/v1/todo').pipe(tap(entities => {
      this.todoStore.set(entities);
    }));
  }

  add(todo: TodoItem) {
    //this.todoStore.add(todo);



     return this.http.post<TodoItem[]>('http://localhost:8080/v1/todo/create', todo).pipe(tap(entities => {
       //this.todoStore.set(entities);
     }));

  }

  update(id: ID, todo: Partial<TodoItem>) {
    this.todoStore.update(id, todo);
  }

  remove(id: ID) {
    this.todoStore.remove(id);
  }

}
