import { Module } from '@nestjs/common';
import { EventsCoreModule } from '@app/core';
import { EventsController } from './events.controller';

@Module({
  imports: [EventsCoreModule],
  controllers: [EventsController],
})
export class EventsModule {}
