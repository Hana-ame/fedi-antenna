// const { db } = require('@vercel/postgres');
import { PrismaClient } from '@prisma/client'
const prisma = new PrismaClient()
import {
  invoices,
  customers,
  revenue,
  users,
} from '../app/lib/placeholder-data.js'
import bcrypt from 'bcrypt';

async function seedUsers() {
  try {
    // Insert data into the "users" table
    const insertedUsers = await Promise.all(
      users.map(async (user) => {
        const hashedPassword = await bcrypt.hash(user.password, 10);
        return prisma.$queryRaw`
          INSERT INTO users (id, name, email, password)
          VALUES (uuid(${user.id}), ${user.name}, ${user.email}, ${hashedPassword})
          ON CONFLICT (id) DO NOTHING;
        `;
      }),
    );

    console.log(`Seeded ${insertedUsers.length} users`);

  } catch (error) {
    console.error('Error seeding users:', error);
    throw error;
  }
}

async function seedInvoices(client) {
  try {
    // Insert data into the "invoices" table
    const insertedInvoices = await Promise.all(
      invoices.map(
        (invoice) => prisma.$queryRaw`
          INSERT INTO invoices (customer_id, amount, status, date)
          VALUES (uuid(${invoice.customer_id}), ${invoice.amount}, ${invoice.status}, date(${invoice.date}))
          ON CONFLICT (id) DO NOTHING;
        `,
      ),
    );

    console.log(`Seeded ${insertedInvoices.length} invoices`);

  } catch (error) {
    console.error('Error seeding invoices:', error);
    throw error;
  }
}

async function seedCustomers(client) {
  try {
    // Insert data into the "customers" table
    const insertedCustomers = await Promise.all(
      customers.map(
        (customer) => prisma.$queryRaw`
          INSERT INTO customers (id, name, email, image_url)
          VALUES (uuid(${customer.id}), ${customer.name}, ${customer.email}, ${customer.image_url})
          ON CONFLICT (id) DO NOTHING;
        `,
      ),
    );

    console.log(`Seeded ${insertedCustomers.length} customers`);

  } catch (error) {
    console.error('Error seeding customers:', error);
    throw error;
  }
}

async function seedRevenue() {
  try {
    // Insert data into the "revenue" table
    const insertedRevenue = await Promise.all(
      revenue.map(
        (rev) => prisma.$queryRaw`
          INSERT INTO revenue (month, revenue)
          VALUES (${rev.month}, ${rev.revenue})
          ON CONFLICT (month) DO NOTHING;
        `,
      ),
    );

    console.log(`Seeded ${insertedRevenue.length} revenue`);

  } catch (error) {
    console.error('Error seeding revenue:', error);
    throw error;
  }
}

async function main() {

  await seedUsers();
  await seedCustomers();
  await seedInvoices();
  await seedRevenue();

  await prisma.$disconnect()
}

main().catch((err) => {
  console.error(
    'An error occurred while attempting to seed the database:',
    err,
  );
});
