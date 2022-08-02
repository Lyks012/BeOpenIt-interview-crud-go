import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { Post } from 'src/app/model/post';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css'],
})
export class PostComponent implements OnInit {
  ngOnInit(): void {}
  @Input() post: Post | undefined;

  @Output() postToDelete: EventEmitter<Post> = new EventEmitter();
  handleDelete() {
    this.postToDelete.emit();
  }
}
