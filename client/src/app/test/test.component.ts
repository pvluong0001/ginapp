import {EventEmitter, Component, Input, OnChanges, OnInit, Output, SimpleChanges} from '@angular/core';

@Component({
  selector: 'app-test',
  templateUrl: './test.component.html',
  styleUrls: ['./test.component.scss']
})
export class TestComponent implements OnInit, OnChanges {
  @Input() progress = 0;
  @Output() submitEvent = new EventEmitter<string>();
  @Input() toggle = true;
  @Output() toggleChange = new EventEmitter<boolean>();

  constructor() {
  }

  ngOnChanges(changes: SimpleChanges): void {
    if ('progress' in changes) {
      if (typeof changes.progress.currentValue !== 'number') {
        const progress = Number(changes.progress.currentValue);
        if (Number.isNaN(progress)) {
          this.progress = 0;
        } else {
          this.progress = progress;
        }
      }
    }
  }

  ngOnInit(): void {
  }

  handleClick(): void {
    this.submitEvent.emit('data from parent');
    this.toggle = !this.toggle;
    this.toggleChange.emit(this.toggle);
  }
}
