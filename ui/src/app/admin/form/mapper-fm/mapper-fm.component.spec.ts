import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MapperFmComponent } from './mapper-fm.component';

describe('MapperFmComponent', () => {
  let component: MapperFmComponent;
  let fixture: ComponentFixture<MapperFmComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MapperFmComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MapperFmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
