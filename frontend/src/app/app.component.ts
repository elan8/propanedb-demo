import { Component } from '@angular/core';
import { createTodo, TodoItem } from './todo/state/todo.model';
import { TodoService } from './todo/state/todo.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {
  title = 'frontend';
  todos: TodoItem[] = [];
  //todo1 =null;

  constructor(private todoService: TodoService) {

    var todo1=createTodo("Todo 1");

    todoService.add(todo1).subscribe( data => {
      console.log("Add: Data="+ JSON.stringify(data) )
     this.todos= data
    })

   todoService.get().subscribe( data => {
     console.log("Get: Data="+ JSON.stringify(data) )
    this.todos= data
   })

  }
}

