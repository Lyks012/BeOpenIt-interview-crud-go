import { Post } from 'src/app/model/post';
import { Category } from 'src/app/model/category';
import { Component, OnInit, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {
  descriptionPost:string = "";
  postCategorieId!: number;

  @Output() createCategorie = new EventEmitter();

  @Output() createProduit = new EventEmitter();

  @Input() categories : Category[] | undefined ;

  handleCreateCategorie(nomCategorie:any){
    this.createCategorie.emit(new Category(0, nomCategorie.value, []));
  }

  handleCreatePost(nomPost:any){
    this.createProduit.emit(new Post(0, nomPost.value, this.descriptionPost, this.postCategorieId, ""))
  }

  constructor() { }

  ngOnInit(): void {
  }

}
