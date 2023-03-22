import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LinkFmComponent } from './link-fm.component';

describe('LinkFmComponent', () => {
  let component: LinkFmComponent;
  let fixture: ComponentFixture<LinkFmComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LinkFmComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LinkFmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
