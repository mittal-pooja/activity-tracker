-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 16, 2020 at 11:11 PM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.2.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `activity_tracker_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `activity_table`
--

CREATE TABLE `activity_table` (
  `id` int(11) NOT NULL,
  `Type` varchar(10) NOT NULL,
  `Label` varchar(100) NOT NULL,
  `Duration` time NOT NULL DEFAULT '00:00:00',
  `Timestamp` timestamp NOT NULL DEFAULT current_timestamp(),
  `Status` varchar(10) NOT NULL DEFAULT 'NOT DONE',
  `Valid` varchar(10) NOT NULL DEFAULT 'FALSE'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `activity_table`
--

INSERT INTO `activity_table` (`id`, `Type`, `Label`, `Duration`, `Timestamp`, `Status`, `Valid`) VALUES
(2, 'Sleep', 'Pooja', '00:00:00', '2020-11-09 20:47:01', 'NOT DONE', 'FALSE'),
(3, 'Play', 'Neha', '00:00:00', '2020-11-09 20:47:01', 'NOT DONE', 'FALSE'),
(13, 'Eat', 'Pooja', '07:33:00', '2020-11-16 08:47:38', 'DONE', 'VALID');

-- --------------------------------------------------------

--
-- Table structure for table `user_table`
--

CREATE TABLE `user_table` (
  `Name` varchar(20) NOT NULL,
  `Email` varchar(50) NOT NULL,
  `Phone` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user_table`
--

INSERT INTO `user_table` (`Name`, `Email`, `Phone`) VALUES
('Neha', 'agarwalneaha@example.com', '1234567894'),
('Pooja', '12345689@example.com', '1234567894');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `activity_table`
--
ALTER TABLE `activity_table`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_activity_label` (`Label`);

--
-- Indexes for table `user_table`
--
ALTER TABLE `user_table`
  ADD PRIMARY KEY (`Name`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `activity_table`
--
ALTER TABLE `activity_table`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `activity_table`
--
ALTER TABLE `activity_table`
  ADD CONSTRAINT `fk_activity_label` FOREIGN KEY (`Label`) REFERENCES `user_table` (`Name`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
