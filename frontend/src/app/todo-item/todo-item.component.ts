import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { TodoItem } from '../todo/state/todo.model';
import {MatCheckboxChange} from '@angular/material/checkbox';

@Component({
  selector: 'app-todo-item',
  templateUrl: './todo-item.component.html',
  styleUrls: ['./todo-item.component.sass']
})
export class TodoItemComponent implements OnInit {

  @Input()
  todo!: TodoItem;

  @Output() completeChange = new EventEmitter<MatCheckboxChange>();

  constructor() { }

  ngOnInit(): void {
  }

}
