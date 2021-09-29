import { Injectable } from '@angular/core';
import { EntityState, EntityStore, StoreConfig } from '@datorama/akita';
import { TodoItem } from './todo.model';

export interface TodoState extends EntityState<TodoItem> {}

@Injectable({ providedIn: 'root' })
@StoreConfig({ name: 'todo' })
export class TodoStore extends EntityStore<TodoState> {

  constructor() {
    super();
  }

}
