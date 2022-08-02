import { Category } from 'src/app/model/category';
import { Component, Input, Output, OnInit, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-filtres',
  templateUrl: './filtres.component.html',
  styleUrls: ['./filtres.component.css']
})
export class FiltresComponent implements OnInit {
  postName: string = ""
  categorySelected?: Category;
  @Input() categories : Category[] | undefined;

  @Output() categoryToFilter: EventEmitter<Category> = new EventEmitter();
  handleFilterByCategory(){
    this.categoryToFilter.emit(this.categorySelected);
  }

  @Output() nameToFind: EventEmitter<string> = new EventEmitter();
  handleFindPostName(){
    this.nameToFind.emit(this.postName)
  }
  constructor() { }

  ngOnInit(): void {
  }

}
