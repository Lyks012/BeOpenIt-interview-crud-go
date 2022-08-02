import { Post } from './post';
export class Category {
    public Id: number;
    public Name:string;
    public Posts:Post[];
  constructor(id: number, name: string, posts: Post[]) {
    this.Id = id;
    this.Name = name;
    this.Posts = posts;
  }
}
