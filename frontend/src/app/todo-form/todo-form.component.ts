import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Subject } from 'rxjs';
import { debounceTime, takeUntil } from 'rxjs/operators';
import { TodoItem } from '../todo-state/todo.model';

@Component({
  selector: 'app-todo-form',
  templateUrl: './todo-form.component.html',
  styleUrls: ['./todo-form.component.scss']
})
export class TodoFormComponent implements OnInit {
  @Output() addTodoEmitter = new EventEmitter<string>();

  title: FormControl = new FormControl;

  //private unsubscribe = new Subject<void>();

  constructor() { }

  ngOnDestroy() {
    //this.unsubscribe.next();
    //this.unsubscribe.complete();
  }

  ngOnInit() {
    this.title = new FormControl();
    // this.title.valueChanges
    //   .pipe(debounceTime(200), takeUntil(this.unsubscribe))
    //   .subscribe(value => this.toDoChange.emit( value ));
  }
  addTodo() {
    this.addTodoEmitter.emit(this.title.value);
    this.title.reset();
    // this.store.dispatch(new AddToDo({
    //   id: Math.random(),
    //   complete: false,
    //   task: this._toDo.task
    // }));

    // var todo1 = createTodo(this.title);
    // this._todoService.add(todo1).subscribe(data => {
    //   console.log("Add: Data=" + JSON.stringify(data))
    // })

    // //update todos
    // this.todoService.getAll().subscribe(response => {
    //   console.log("Get: Data=" + JSON.stringify(response))
    //   //this.todos= data
    //   let todoItems = response?.body as TodoItems;
    //   let todos = todoItems.items;

    // })

  }

}
