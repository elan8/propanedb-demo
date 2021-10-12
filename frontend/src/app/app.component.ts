import { Component } from '@angular/core';
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
  //todos: HashMap<TodoItem> | undefined ;
  //todo1 =null;

  constructor(private todoService: TodoService , private todoQuery : TodoQuery) {

    this.todoQuery.select(state => state.todos).subscribe(
      todos => {
        this.todos = todos;
        //console.log("Todos="+todos);
        // this.maxNumberOfServices = user.accountTypeDefinition.numberOfServices;
        // this.maxNumberOfRPC = user.accountTypeDefinition.numberOfRPC;
        // this.persistenceAvailable=user.accountTypeDefinition.persistence;
        // this.freeAccount = user.accountType=="FREE";
        //this.numberOfProjects= this.CountNumberOfProjects(projects);
        // console.log("Get token:"+ value);
        //this.httpHeaders = new HttpHeaders().set('Authorization', this.token);
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
}

