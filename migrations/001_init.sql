-- Extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Departments jadvali
CREATE TABLE IF NOT EXISTS departments (
                                           id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- Employees jadvali
CREATE TABLE IF NOT EXISTS employees (
                                         id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name     VARCHAR(150) NOT NULL,
    email         VARCHAR(150) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role          VARCHAR(20)  NOT NULL DEFAULT 'employee'
    CHECK (role IN ('employee', 'hr', 'admin')),
    department_id UUID NOT NULL REFERENCES departments(id),
    created_at    TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- Geofence zones jadvali
CREATE TABLE IF NOT EXISTS geofence_zones (
                                              id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR(100) NOT NULL,
    latitude   DOUBLE PRECISION NOT NULL,
    longitude  DOUBLE PRECISION NOT NULL,
    radius_m   DOUBLE PRECISION NOT NULL,
    is_active  BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- Attendance jadvali
CREATE TABLE IF NOT EXISTS attendance (
                                          id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id UUID NOT NULL REFERENCES employees(id),
    check_in    TIMESTAMP,
    check_out   TIMESTAMP,
    status      VARCHAR(20) NOT NULL DEFAULT 'present'
    CHECK (status IN ('present', 'absent', 'late')),
    latitude    DOUBLE PRECISION NOT NULL DEFAULT 0,
    longitude   DOUBLE PRECISION NOT NULL DEFAULT 0,
    date        DATE NOT NULL DEFAULT CURRENT_DATE,
    UNIQUE (employee_id, date)
    );

-- Leave requests jadvali
CREATE TABLE IF NOT EXISTS leave_requests (
                                              id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id UUID NOT NULL REFERENCES employees(id),
    leave_type  VARCHAR(20) NOT NULL
    CHECK (leave_type IN ('annual', 'sick', 'unpaid')),
    start_date  DATE NOT NULL,
    end_date    DATE NOT NULL,
    reason      TEXT NOT NULL,
    status      VARCHAR(20) NOT NULL DEFAULT 'pending'
    CHECK (status IN ('pending', 'approved', 'rejected')),
    reviewed_by UUID REFERENCES employees(id),
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- Indexlar (qidiruv tezligi uchun)
CREATE INDEX IF NOT EXISTS idx_attendance_employee ON attendance(employee_id);
CREATE INDEX IF NOT EXISTS idx_attendance_date ON attendance(date);
CREATE INDEX IF NOT EXISTS idx_leave_employee ON leave_requests(employee_id);
CREATE INDEX IF NOT EXISTS idx_leave_status ON leave_requests(status);

-- Test uchun boshlang'ich ma'lumotlar
INSERT INTO departments (id, name) VALUES
                                       ('11111111-1111-1111-1111-111111111111', 'IT Department'),
                                       ('22222222-2222-2222-2222-222222222222', 'HR Department'),
                                       ('33333333-3333-3333-3333-333333333333', 'Finance Department');

-- Ofis geofence zone (Toshkent markazi, 200 metr radius)
INSERT INTO geofence_zones (name, latitude, longitude, radius_m, is_active) VALUES
    ('Bosh ofis', 41.2995, 69.2401, 200.0, true);