import { Component, EventEmitter, Input, Output } from '@angular/core';
import { MatCheckboxChange } from '@angular/material/checkbox';
import { HashMap } from '@datorama/akita';
import { createTodo, TodoItem } from './todo/state/todo.model';
import { TodoQuery } from './todo/state/todo.query';
import { TodoService } from './todo/state/todo.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {
  title = 'frontend';
  todos: TodoItem[] = [];

  // @Input()
  // todos: TodoItem[] = [];

  @Output() todoChange = new EventEmitter<TodoItem>();

  constructor(private todoService: TodoService , private todoQuery : TodoQuery) {

    this.todoQuery.select(state => state.todos).subscribe(
      todos => {
        this.todos = todos;
      }
    )
    var todo1=createTodo("Todo 1");
    todoService.add(todo1).subscribe( data => {
      console.log("Add: Data="+ JSON.stringify(data) )
     //this.todos= data
    })


    // var todo2=createTodo("Todo 2");

    // todoService.add(todo2).subscribe( data => {
    //   console.log("Add: Data="+ JSON.stringify(data) )
    //  //this.todos= data
    // })
  
   todoService.getAll().subscribe( data => {
     //console.log("Get: Data="+ JSON.stringify(data) )
     //this.todos= data
   })

  }


  onCompleteChange(todo: TodoItem, change: MatCheckboxChange) {
    this.todoChange.emit({
      ...todo,
      completed: change.checked
    });
  }

}

