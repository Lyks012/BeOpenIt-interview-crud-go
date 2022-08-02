import { CategoryService } from './../../services/category.service';
import { PostsService } from './../../services/posts.service';
import { Component, OnInit } from '@angular/core';
import { Post } from 'src/app/model/post';
import { Category } from 'src/app/model/category';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css'],
})
export class PostsComponent implements OnInit {
  posts: Post[] = [];
  categories: Category[] = [];

  displayedCategoriesWithPosts: Category[] = [];

  constructor(
    private postService: PostsService,
    private categoryService: CategoryService
  ) {}

  ngOnInit(): void {
    this.categoryService.getAllCategories().subscribe((res) => {
      this.categories = res.Data;
      this.fillPostsOfCategory();
    });
  }
  private fillPostsOfCategory() {
    for (let i = 0; i < this.categories.length; i++) {
      let category = this.categories[i];
      this.postService.getByCategoryId(category.Id).subscribe((res) => {
        const posts = res;
        category.Posts = posts;
        this.displayedCategoriesWithPosts.push(category);
      });
    }
  }

  filterByCategory(categoryToFilter: Category) {
    if (typeof categoryToFilter !== 'string')
      this.displayedCategoriesWithPosts = this.categories.filter((category) => {
        return category == categoryToFilter;
      });
    else this.displayedCategoriesWithPosts = this.categories;
  }

  findByPostName(name: string) {
    let found: boolean = false;

    this.categories.forEach((categoryWithPosts) => {
      if (
        categoryWithPosts.Posts !== null &&
        categoryWithPosts.Posts.length > 0
      ) {
        categoryWithPosts.Posts.forEach((post) => {
          // update the displayed categories with posts when found
          if (post.Title == name) {
            this.displayedCategoriesWithPosts = [];
            let posts = [];
            posts.push(post);

            this.displayedCategoriesWithPosts.push(
              new Category(categoryWithPosts.Id, categoryWithPosts.Name, posts)
            );
            found = true;
            return;
          }
        });
      }
    });
    if (!found) this.displayedCategoriesWithPosts = [];
  }

  deletePost(postToDelete: Post) {
    this.postService
      .deletePost(postToDelete.Id)
      .subscribe((res) => console.log(res));

      // update the data locally 
    this.displayedCategoriesWithPosts.forEach((category) => {
      if (category.Posts !== null && category.Posts.length > 0)
        category.Posts.forEach((post) => {
          if (post == postToDelete) {
            category.Posts.splice(category.Posts.indexOf(post), 1);
            return;
          }
        });
    });
  }

  createOneCategory(category: Category) {
    this.categoryService
      .createCategory(category)
      .subscribe((res) => console.log(res));
  }

  createOneProduit(post: Post) {
    this.postService.createPost(post).subscribe((res) => console.log(res));
  }
}
