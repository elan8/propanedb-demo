import { Component, EventEmitter, Input, Output } from '@angular/core';
import { MatCheckboxChange } from '@angular/material/checkbox';
import { HashMap } from '@datorama/akita';
import { createTodo, TodoItem, TodoItems } from './todo-state/todo.model';
import { TodoQuery } from './todo-state/todo.query';
import { TodoService } from './todo-state/todo.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  //title = 'frontend';
  todos: TodoItem[] = [];
  title: string = "";
  _todoService: TodoService;


  @Output() todoChange = new EventEmitter<TodoItem>();

  constructor(private todoService: TodoService, private todoQuery: TodoQuery) {

    this._todoService = todoService;
    this.todoQuery.select(state => state.todos).subscribe(
      todos => {
        this.todos = todos;
      }
    )

    this.getTodos();


  }

  getTodos() {
    //update todos
    this.todoService.getAll().subscribe(response => {
      console.log("Get: Data=" + JSON.stringify(response))
      //this.todos= data
      let todoItems = response?.body as TodoItems;
      let todos = todoItems.items;

    })
  }



  addTodo(title: string) {
    // this.store.dispatch(new AddToDo({
    //   id: Math.random(),
    //   complete: false,
    //   task: this._toDo.task
    // }));

    var todo1 = createTodo(title);
    this._todoService.add(todo1).subscribe(data => {
      console.log("Add: Data=" + JSON.stringify(data))
    })

    this.getTodos();
    // //update todos
    // this.todoService.getAll().subscribe(response => {
    //   console.log("Get: Data=" + JSON.stringify(response))
    //   //this.todos= data
    //   let todoItems = response?.body as TodoItems;
    //   let todos = todoItems.items;

    // })

  }

  // onAddToDoChange(title: string) {
  //   this.title = title;
  // }

  onCompleteChange(todo: TodoItem, change: MatCheckboxChange) {

    this._todoService.update(todo).subscribe(data => {
      console.log("Update todo: " + JSON.stringify(data))
    })

    // this.todoChange.emit({
    //   ...todo,
    //   completed: change.checked
    // });
  }

}

