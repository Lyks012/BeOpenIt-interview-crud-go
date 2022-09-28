-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3307
-- Generation Time: Aug 02, 2022 at 12:03 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crud-go`
--

-- --------------------------------------------------------

--
-- Table structure for table `category`
--

CREATE TABLE `category` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `category`
--

INSERT INTO `category` (`id`, `name`) VALUES
(27, 'BeOpen'),
(28, 'CNO');

-- --------------------------------------------------------

--
-- Table structure for table `post`
--

CREATE TABLE `post` (
  `id` int(11) NOT NULL,
  `title` varchar(50) NOT NULL,
  `description` varchar(255) NOT NULL,
  `id_category` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `post`
--

INSERT INTO `post` (`id`, `title`, `description`, `id_category`) VALUES
(32, 'docker', 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Ipsam, excepturi minima distinctio officia expedita iure ex ipsa optio in placeat eum voluptas impedit, itaque consectetur mollitia laudantium ducimus esse consequatur!', 27),
(33, 'kubernetes', 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Ipsam, excepturi minima distinctio officia expedita iure ex ipsa optio in placeat eum voluptas impedit, itaque consectetur mollitia laudantium ducimus esse consequatur!', 27),
(34, 'Kubernetes-cli', 'Lorem ipsum, dolor sit amet consectetur adipisicing elit. Ipsam, excepturi minima distinctio officia expedita iure ex ipsa optio in placeat eum voluptas impedit, itaque consectetur mollitia laudantium ducimus esse consequatur!', 28),
(35, 'docker_cli', '\r\nLorem ipsum, dolor sit amet consectetur adipisicing elit. Ipsam, excepturi minima distinctio officia expedita iure ex ipsa optio in placeat eum voluptas impedit, itaque consectetur mollitia laudantium ducimus esse consequatur!', 28);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `post`
--
ALTER TABLE `post`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_category` (`id_category`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `category`
--
ALTER TABLE `category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT for table `post`
--
ALTER TABLE `post`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `post`
--
ALTER TABLE `post`
  ADD CONSTRAINT `post_ibfk_1` FOREIGN KEY (`id_category`) REFERENCES `category` (`id`);